/**
 * @module modules/plot-summary-sk
 * @description <h2><code>plot-summary-sk</code></h2>
 *
 * @evt summary_selected - Event produced when the user selects a section on the summary.
 *
 * @attr
 *
 * @example
 */
import { html } from 'lit-html';
import * as d3Scale from 'd3-scale';
import { load } from '@google-web-components/google-chart/loader';
import { define } from '../../../elements-sk/modules/define';
import { ElementSk } from '../../../infra-sk/modules/ElementSk';
import { MousePosition, Point } from '../plot-simple-sk/plot-simple-sk';
import '@google-web-components/google-chart/';
import { ChartData, DrawSummaryChart } from '../common/plot-builder';

const ZOOM_RECT_COLOR = '#0007';

// Describes the zoom in terms of x-axis source values.
export type ZoomRange = [number, number] | null;
export interface PlotSummarySkSelectionEventDetails {
  start: number;
  end: number;
  valueStart: number | Date;
  valueEnd: number | Date;
}

export class PlotSummarySk extends ElementSk {
  constructor() {
    super(PlotSummarySk.template);
  }

  private overlayCtx: CanvasRenderingContext2D | null = null;

  /** The window.devicePixelRatio. */
  private scale: number = 1.0;

  // This contains a mapping of the coordinates on the summary bar to
  // the corresponding values. This helps us translate the position of
  // the selected area to the corresponding values that they represent.
  private valuesRangeCommit: d3Scale.ScaleLinear<number, number> | null = null;

  private commitsStart: number = 0;

  private commitsEnd: number = 0;

  private valuesRangeDate: d3Scale.ScaleTime<number, number> | null = null;

  private dateStart: Date = new Date();

  private dateEnd: Date = new Date();

  private isCommitScale: boolean = false;

  // Array denoting the start and end points of the current selection.
  private selectionRange: ZoomRange = null;

  // Array denoting the diffs between the currently clicked point and
  // the start and end points of the current selection. This is used
  // to maintain the size of the selection when user drags it left/right.
  private lockedSelectionDiffs: [number, number] | null = null;

  // Denotes if the user is currently making a selection.
  private isCurrentlySelecting: boolean = false;

  // Tracks the mouse position.
  private currentMousePosition: MousePosition | null = null;

  // The div element that will host the plot on the summary.
  private plotElement: HTMLElement | null = null;

  // The canvas element that is used for the selection overlay.
  private overlayCanvas: HTMLCanvasElement | null = null;

  // Keeps a track of the current chart data being displayed.
  private currentChartData: ChartData | null = null;

  private static template = (ele: PlotSummarySk) => html`
    <div class="border">
      <div
        id="plot"
        class="plot"
        width=${ele.width * window.devicePixelRatio}
        height=${ele.height * window.devicePixelRatio}
        style="transform-origin: 0 0; transform: scale(${1 /
        window.devicePixelRatio});"></div>
      <canvas
        id="overlay"
        class="overlay"
        width=${ele.width * window.devicePixelRatio}
        height=${ele.height * window.devicePixelRatio}
        style="transform-origin: 0 0; transform: scale(${1 /
        window.devicePixelRatio});"></canvas>
    </div>
  `;

  async connectedCallback(): Promise<void> {
    super.connectedCallback();
    this.render();

    window.requestAnimationFrame(this.raf.bind(this));
  }

  render(): void {
    this._render();
    this.scale = window.devicePixelRatio;
    this.plotElement = this.querySelector<HTMLElement>('#plot')!;
    this.overlayCanvas = this.querySelector<HTMLCanvasElement>('#overlay')!;
    this.overlayCtx = this.overlayCanvas.getContext('2d');

    // globalAlpha denotes the transparency of the fill. Setting this to 50%
    // allows us to show the highlight plus the portion of the plot highlighted.
    this.overlayCtx!.globalAlpha = 0.5;

    this.addEventListeners();
    this.drawSummaryRect();
  }

  private async _renderChart() {
    if (this.isCommitScale) {
      this.commitsStart = this.currentChartData!.data[0].x as number;
      this.commitsEnd = this.currentChartData!.data[
        this.currentChartData!.data.length - 1
      ].x as number;
      this.valuesRangeCommit = d3Scale
        .scaleLinear()
        .domain([0, this.width])
        .range([this.commitsStart, this.commitsEnd]);
    } else {
      this.dateStart = this.currentChartData!.data[0].x as Date;
      this.dateEnd = this.currentChartData!.data[
        this.currentChartData!.data.length - 1
      ].x as Date;
      this.valuesRangeDate = d3Scale
        .scaleTime()
        .domain([0, this.width])
        .range([this.dateStart.getTime(), this.dateEnd.getTime()]);
    }

    await load({ packages: ['corechart'] });
    DrawSummaryChart(
      this.plotElement!,
      this.currentChartData!,
      this.width,
      this.height,
      getComputedStyle(this)
    );
  }

  // Select the provided range on the plot-summary.
  public Select(valueStart: number | Date, valuesEnd: number | Date) {
    if (this.isCommitScale) {
      this.selectionRange = [
        this.valuesRangeCommit!(valueStart),
        this.valuesRangeCommit!(valuesEnd),
      ];
    } else {
      this.selectionRange = [
        this.valuesRangeDate!(valueStart),
        this.valuesRangeDate!(valuesEnd),
      ];
    }

    // Dispatch the summary selection event.
    this.summarySelected();
  }

  // Display the chart data on the plot.
  public async DisplayChartData(chartData: ChartData, isCommitScale: boolean) {
    if (chartData.data?.length > 0) {
      this.currentChartData = chartData;
      this.isCommitScale = isCommitScale;
      this._renderChart();
    }
  }

  // Clear the current selection.
  private clearSelection(): void {
    this.overlayCtx!.clearRect(
      0,
      0,
      this.overlayCanvas!.width,
      this.overlayCanvas!.height
    );
  }

  // Listener for the mouse down event.
  private mouseDownListener(e: MouseEvent) {
    e.preventDefault();
    const point = this.eventToCanvasPt(e);
    this.currentMousePosition = {
      clientX: point.x,
      clientY: point.y,
    };

    if (this.inSelectedArea(point)) {
      // Implement the drag functionality to drag the selected area around.
      // This means the user is dragging the current selection around.
      // Let's get the left and right diffs between the current mouse position
      // and the ends of the selection range. As the mouse moves, we will update
      // the selection range so that these diffs remain constant.
      const leftDiff =
        this.currentMousePosition.clientX - this.selectionRange![0];
      const rightDiff =
        this.selectionRange![1] - this.currentMousePosition.clientX;
      this.lockedSelectionDiffs = [leftDiff, rightDiff];
    } else {
      // User is starting a new selection.
      this.selectionRange = [point.x, point.x + 0.1];
      this.isCurrentlySelecting = true;
    }
  }

  // Listener for the mouse up event.
  private mouseUpListener(e: MouseEvent) {
    // Releasing the mouse means selection/dragging is done.
    this.isCurrentlySelecting = false;
    this.lockedSelectionDiffs = null;
    this.summarySelected();
  }

  private summarySelected() {
    if (this.selectionRange !== null) {
      let start: number | Date;
      let end: number | Date;
      if (this.isCommitScale) {
        start = this.valuesRangeCommit!(this.selectionRange[0]);
        end = this.valuesRangeCommit!(this.selectionRange[1]);
      } else {
        start = new Date(this.valuesRangeDate!(this.selectionRange[0]));
        end = new Date(this.valuesRangeDate!(this.selectionRange[1]));
      }
      this.dispatchEvent(
        new CustomEvent<PlotSummarySkSelectionEventDetails>(
          'summary_selected',
          {
            detail: {
              start: this.selectionRange[0],
              end: this.selectionRange[1],
              valueStart: start,
              valueEnd: end,
            },
            bubbles: true,
          }
        )
      );
    }
  }

  // Listener for the mouse move event.
  private mouseMoveListener(e: MouseEvent) {
    // Keep track of the user's mouse movements.
    const point = this.eventToCanvasPt(e);
    this.currentMousePosition = {
      clientX: point.x,
      clientY: point.y,
    };
  }

  // Add all the event listeners.
  private addEventListeners(): void {
    // If the user toggles the theme to/from darkmode then redraw.
    document.addEventListener('theme-chooser-toggle', () => {
      this.render();
      this._renderChart();
    });
    this.addEventListener('mousedown', this.mouseDownListener);
    this.addEventListener('mouseup', this.mouseUpListener);
    this.addEventListener('mouseleave', this.mouseUpListener);
    this.addEventListener('mousemove', this.mouseMoveListener);
  }

  // Draw the summary rectangle outline.
  private drawSummaryRect(): void {
    const style = getComputedStyle(this);
    this.overlayCtx!.lineWidth = 3.0;
    this.overlayCtx!.strokeStyle = style.color;
    this.overlayCtx!.strokeRect(0, 0, this.width, this.height);
    this.overlayCtx!.save();
  }

  /** Mirrors the width attribute. */
  get width(): number {
    return +(this.getAttribute('width') || '0');
  }

  set width(val: number) {
    this.setAttribute('width', val.toString());
  }

  /** Mirrors the height attribute. */
  get height(): number {
    return +(this.getAttribute('height') || '0');
  }

  set height(val: number) {
    this.setAttribute('height', val.toString());
  }

  /** Mirrors the highlight_color attribute. */
  get highlightColor(): string {
    return this.getAttribute('highlight_color') || ZOOM_RECT_COLOR;
  }

  set highlightColor(val: string) {
    this.setAttribute('highlight_color', val);
  }

  // Converts an event to a specific point
  private eventToCanvasPt(e: MousePosition) {
    const clientRect = this.overlayCtx!.canvas.getBoundingClientRect();
    return {
      x: (e.clientX - clientRect.left) * this.scale,
      y: (e.clientY - clientRect.top) * this.scale,
    };
  }

  // Draws the selection area.
  private drawSelection(): void {
    this.clearSelection();
    if (this.selectionRange !== null) {
      // Draw left line.
      const startx = this.selectionRange[0];
      this.drawVerticalLineAtPosition(startx);
      // Draw the right line.
      const endx = this.selectionRange[1];
      this.drawVerticalLineAtPosition(endx);

      // Shade the selected section.
      this.overlayCtx!.beginPath();
      this.overlayCtx!.fillStyle = this.highlightColor;
      this.overlayCtx!.rect(startx, 0, endx - startx, this.height);
      this.overlayCtx!.fill();
    }
  }

  // Draw a vertical line at the given position
  private drawVerticalLineAtPosition(x: number) {
    // Start at the bottom of the rectangle
    this.overlayCtx!.moveTo(x, 0);
    this.overlayCtx!.lineTo(x, this.height);
    this.overlayCtx!.stroke();
  }

  // Handles the animation depending on the current state.
  private raf() {
    // Always queue up our next raf first.
    window.requestAnimationFrame(() => this.raf());

    // Exit early if there is no ongoing activity.
    if (this.currentMousePosition === null) {
      return;
    }

    if (this.isCurrentlySelecting) {
      // If the user is currently selecting an area, update the selection range
      // array based on the current mouse position.
      const startx = this.selectionRange![0];
      let endx = this.currentMousePosition.clientX;
      if (endx < startx) {
        endx = startx;
      }
      this.selectionRange![1] = endx;
    } else if (this.lockedSelectionDiffs !== null) {
      // User is dragging the current selection around.
      // Update the selectionRange with the values that keeps these
      // diffs constant wrt the current mouse position. Also account for the fact
      // that the user may drag the selected area beyond the bounds of the rendered rect.
      this.selectionRange![0] = Math.max(
        0,
        this.currentMousePosition.clientX - this.lockedSelectionDiffs![0]
      );
      this.selectionRange![1] = Math.min(
        this.width,
        this.currentMousePosition.clientX + this.lockedSelectionDiffs![1]
      );
    }

    this.drawSelection();
  }

  // Checks if a point is inside the selected area.
  private inSelectedArea(pt: Point): boolean {
    if (this.selectionRange === null) {
      return false;
    }

    return pt.x >= this.selectionRange[0] && pt.x < this.selectionRange[1];
  }
}

define('plot-summary-sk', PlotSummarySk);
