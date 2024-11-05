import './index';
import fetchMock from 'fetch-mock';
import { JSONSourceSk } from './json-source-sk';
import '../../../elements-sk/modules/error-toast-sk';
import { CommitNumber } from '../json';

fetchMock.post('/_/details/', () => ({ Hello: 'world!' }));

window.customElements.whenDefined('json-source-sk').then(() => {
  const sources = document.querySelectorAll<JSONSourceSk>('json-source-sk')!;
  sources.forEach((source) => {
    source.traceid = ',foo=bar,';
    source.cid = CommitNumber(12);
    source.querySelector<HTMLButtonElement>('#controls button')!.click();
  });
});
