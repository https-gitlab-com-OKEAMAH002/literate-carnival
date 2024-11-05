import { GetTasksResponse } from '../json';

export const tasksResult0: GetTasksResponse = {
  data: [
    {
      ts_added: 20200723222929,
      ts_started: 20200723222930,
      ts_completed: 20200724050154,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumPerf-5094\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d969d43a4521410',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2316326/2 ([WIP] Track dynamically added scripts with no src in AdTracker)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/cf90a6c04958ad2ec0d41bda844e9008654b60e6.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/alexmt-ChromiumPerf-5094/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5094-nopatch/consolidated_outputs/alexmt-ChromiumPerf-5094-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5094-withpatch/consolidated_outputs/alexmt-ChromiumPerf-5094-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200723192204,
      ts_started: 20200723192207,
      ts_completed: 20200724003152,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumPerf-5091\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d95f1b4a3891710',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --max-pages-per-bot=10',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2316326/1 ([WIP] Track dynamically added scripts with no src in AdTracker)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/8e2c009191f9ac50e59b8c896f475e2a4c555b59.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/alexmt-ChromiumPerf-5091/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5091-nopatch/consolidated_outputs/alexmt-ChromiumPerf-5091-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5091-withpatch/consolidated_outputs/alexmt-ChromiumPerf-5091-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200723135735,
      ts_started: 20200723135738,
      ts_completed: 20200723203553,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:backer-ChromiumPerf-5084\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d94c8a1d384a910',
      benchmark: 'rendering.cluster_telemetry',
      platform: 'Android',
      run_on_gce: false,
      page_sets: 'Mobile10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch:
        '--enable-features=UseSkiaRenderer --use-gl=egl --disable-features=Vulkan',
      browser_args_with_patch:
        '--enable-features=UseSkiaRenderer --use-gl=angle --use-cmd-decoder=passthrough --disable-features=Vulkan',
      description: 'Testing ANGLE overhead',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/220e34417681c50b03443f1ae74d4d2f12b7b20e.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/220e34417681c50b03443f1ae74d4d2f12b7b20e.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/backer-ChromiumPerf-5084/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/backer-ChromiumPerf-5084-nopatch/consolidated_outputs/backer-ChromiumPerf-5084-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/backer-ChromiumPerf-5084-withpatch/consolidated_outputs/backer-ChromiumPerf-5084-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200723044620,
      ts_started: 20200723044623,
      ts_completed: 20200724005552,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:sauski-ChromiumPerf-5082\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d92cff197daa510',
      benchmark: 'loading.cluster_telemetry',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch:
        '--enable-features=ClientStorageAccessContextAuditing',
      description: 'Client Storage API Access Context Auditing - Cookies',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/sauski-ChromiumPerf-5082/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/sauski-ChromiumPerf-5082-nopatch/consolidated_outputs/sauski-ChromiumPerf-5082-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/sauski-ChromiumPerf-5082-withpatch/consolidated_outputs/sauski-ChromiumPerf-5082-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'avg',
    },
    {
      ts_added: 20200723002847,
      ts_started: 20200723002848,
      ts_completed: 20200723052634,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:backer-ChromiumPerf-5081\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d91e420c1ef9610',
      benchmark: 'rendering.cluster_telemetry',
      platform: 'Android',
      run_on_gce: false,
      page_sets: 'Mobile10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '--enable-features=UseSkiaRenderer',
      browser_args_with_patch:
        '--enable-features=UseSkiaRenderer --use-gl=angle --use-cmd-decoder=passthrough',
      description: 'Testing ANGLE overhead',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/8a9ada38ecc1413abc7180faf960b9c976dc139f.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/8a9ada38ecc1413abc7180faf960b9c976dc139f.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/backer-ChromiumPerf-5081/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/backer-ChromiumPerf-5081-nopatch/consolidated_outputs/backer-ChromiumPerf-5081-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/backer-ChromiumPerf-5081-withpatch/consolidated_outputs/backer-ChromiumPerf-5081-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200721083934,
      ts_started: 20200721083937,
      ts_completed: 20200721133734,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 7,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:sadrul-ChromiumPerf-5074\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d8958c2f8927b10',
      benchmark: 'rendering.cluster_telemetry',
      platform: 'Android',
      run_on_gce: false,
      page_sets: 'Mobile10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'https://bugs.chromium.org/p/chromium/issues/detail?id=968412',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/sadrul-ChromiumPerf-5074/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/sadrul-ChromiumPerf-5074-nopatch/consolidated_outputs/sadrul-ChromiumPerf-5074-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/sadrul-ChromiumPerf-5074-withpatch/consolidated_outputs/sadrul-ChromiumPerf-5074-withpatch.output',
      chromium_hash: '',
      cc_list: ['someone@google.com', 'someone@google.com'],
      task_priority: 100,
      group_name: 'ClusteringTop10K',
      value_column_name: '',
    },
    {
      ts_added: 20200715215236,
      ts_started: 20200715215238,
      ts_completed: 20200716080950,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:teresakang-ChromiumPerf-5065\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d6d48a186847610',
      benchmark: 'rendering.cluster_telemetry',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 10,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=10 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description: 'Testing willReadFrequently',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/e4ed65d9ff42dc58ad00a776310890c76de94277.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/teresakang-ChromiumPerf-5065/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/teresakang-ChromiumPerf-5065-nopatch/consolidated_outputs/teresakang-ChromiumPerf-5065-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/teresakang-ChromiumPerf-5065-withpatch/consolidated_outputs/teresakang-ChromiumPerf-5065-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'avg',
    },
    {
      ts_added: 20200715181424,
      ts_started: 20200715181426,
      ts_completed: 20200715224750,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:cammie-ChromiumPerf-5063\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d6c80dd15c79c10',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Android',
      run_on_gce: false,
      page_sets: 'Mobile10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --additional-histograms=Blink.Paint.UpdateTime,PageLoad.Clients.Ads.FrameCounts.AdFrames.PerFrame.CreativeOriginStatusWithThrottling',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2296293/4 (Perf: NotifyPaintTimingChanged for FEtP only 1 in N times to speed up.)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/a69a1805423e629696d63976e71c64ffc5ddffea.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/cammie-ChromiumPerf-5063/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5063-nopatch/consolidated_outputs/cammie-ChromiumPerf-5063-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5063-withpatch/consolidated_outputs/cammie-ChromiumPerf-5063-withpatch.output',
      chromium_hash: '',
      cc_list: ['cammie@chromium.org'],
      task_priority: 100,
      group_name: '1100404',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200714193806,
      ts_started: 20200714193807,
      ts_completed: 20200715000950,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:cammie-ChromiumPerf-5062\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d67a71f71ba7510',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Android',
      run_on_gce: false,
      page_sets: 'Mobile10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --additional-histograms=Blink.Paint.UpdateTime,PageLoad.Clients.Ads.FrameCounts.AdFrames.PerFrame.CreativeOriginStatusWithThrottling',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2296293/3 (Perf: NotifyPaintTimingChanged for FEtP only 1 in N times to speed up.)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/6dc9bd43d71e9ff3d488861a6ca630a73cbc2d52.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/cammie-ChromiumPerf-5062/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5062-nopatch/consolidated_outputs/cammie-ChromiumPerf-5062-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5062-withpatch/consolidated_outputs/cammie-ChromiumPerf-5062-withpatch.output',
      chromium_hash: '',
      cc_list: ['cammie@chromium.org'],
      task_priority: 100,
      group_name: '1100404',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200714184410,
      ts_started: 20200714184412,
      ts_completed: 20200714235950,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:wangxianzhu-ChromiumPerf-5061\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d6775c135325910',
      benchmark: 'rasterize_and_record_micro_ct',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2297698/1 (Revert "Use PhysicalOffset in legacy inline box painters")',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/107799931181b2f5b13d132af8112c098ca9aa88.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/wangxianzhu-ChromiumPerf-5061/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/wangxianzhu-ChromiumPerf-5061-nopatch/consolidated_outputs/wangxianzhu-ChromiumPerf-5061-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/wangxianzhu-ChromiumPerf-5061-withpatch/consolidated_outputs/wangxianzhu-ChromiumPerf-5061-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'avg',
    },
  ],
  ids: [5094, 5091, 5084, 5082, 5081, 5074, 5065, 5063, 5062, 5061],
  pagination: { offset: 0, size: 10, total: 1827 },
  permissions: [
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
  ],
};
export const tasksResult1: GetTasksResponse = {
  data: [
    {
      ts_added: 20200715215236,
      ts_started: 20200715215238,
      ts_completed: 20200716080950,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:teresakang-ChromiumPerf-5065\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d6d48a186847610',
      benchmark: 'rendering.cluster_telemetry',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 10,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=10 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description: 'Testing willReadFrequently',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/e4ed65d9ff42dc58ad00a776310890c76de94277.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/teresakang-ChromiumPerf-5065/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/teresakang-ChromiumPerf-5065-nopatch/consolidated_outputs/teresakang-ChromiumPerf-5065-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/teresakang-ChromiumPerf-5065-withpatch/consolidated_outputs/teresakang-ChromiumPerf-5065-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'avg',
    },
    {
      ts_added: 20200715181424,
      ts_started: 20200715181426,
      ts_completed: 20200715224750,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:cammie-ChromiumPerf-5063\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d6c80dd15c79c10',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Android',
      run_on_gce: false,
      page_sets: 'Mobile10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --additional-histograms=Blink.Paint.UpdateTime,PageLoad.Clients.Ads.FrameCounts.AdFrames.PerFrame.CreativeOriginStatusWithThrottling',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2296293/4 (Perf: NotifyPaintTimingChanged for FEtP only 1 in N times to speed up.)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/a69a1805423e629696d63976e71c64ffc5ddffea.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/cammie-ChromiumPerf-5063/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5063-nopatch/consolidated_outputs/cammie-ChromiumPerf-5063-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5063-withpatch/consolidated_outputs/cammie-ChromiumPerf-5063-withpatch.output',
      chromium_hash: '',
      cc_list: ['cammie@chromium.org'],
      task_priority: 100,
      group_name: '1100404',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200714193806,
      ts_started: 20200714193807,
      ts_completed: 20200715000950,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:cammie-ChromiumPerf-5062\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d67a71f71ba7510',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Android',
      run_on_gce: false,
      page_sets: 'Mobile10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --additional-histograms=Blink.Paint.UpdateTime,PageLoad.Clients.Ads.FrameCounts.AdFrames.PerFrame.CreativeOriginStatusWithThrottling',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2296293/3 (Perf: NotifyPaintTimingChanged for FEtP only 1 in N times to speed up.)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/6dc9bd43d71e9ff3d488861a6ca630a73cbc2d52.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/cammie-ChromiumPerf-5062/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5062-nopatch/consolidated_outputs/cammie-ChromiumPerf-5062-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/cammie-ChromiumPerf-5062-withpatch/consolidated_outputs/cammie-ChromiumPerf-5062-withpatch.output',
      chromium_hash: '',
      cc_list: ['cammie@chromium.org'],
      task_priority: 100,
      group_name: '1100404',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200714184410,
      ts_started: 20200714184412,
      ts_completed: 20200714235950,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:wangxianzhu-ChromiumPerf-5061\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d6775c135325910',
      benchmark: 'rasterize_and_record_micro_ct',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2297698/1 (Revert "Use PhysicalOffset in legacy inline box painters")',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/107799931181b2f5b13d132af8112c098ca9aa88.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/wangxianzhu-ChromiumPerf-5061/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/wangxianzhu-ChromiumPerf-5061-nopatch/consolidated_outputs/wangxianzhu-ChromiumPerf-5061-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/wangxianzhu-ChromiumPerf-5061-withpatch/consolidated_outputs/wangxianzhu-ChromiumPerf-5061-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'avg',
    },
  ],
  ids: [5094, 5091, 5084, 5082],
  pagination: { offset: 10, size: 10, total: 1827 },
  permissions: [
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
  ],
};
export const tasksResult2: GetTasksResponse = {
  data: [
    {
      ts_added: 20200723222929,
      ts_started: 20200723222930,
      ts_completed: 20200724050154,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumPerf-5094\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d969d43a4521410',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2316326/2 ([WIP] Track dynamically added scripts with no src in AdTracker)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/cf90a6c04958ad2ec0d41bda844e9008654b60e6.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/alexmt-ChromiumPerf-5094/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5094-nopatch/consolidated_outputs/alexmt-ChromiumPerf-5094-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5094-withpatch/consolidated_outputs/alexmt-ChromiumPerf-5094-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'sum',
    },
    {
      ts_added: 20200723192204,
      ts_started: 20200723192207,
      ts_completed: 20200724003152,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumPerf-5091\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d95f1b4a3891710',
      benchmark: 'ad_tagging.cluster_telemetry',
      platform: 'Linux',
      run_on_gce: false,
      page_sets: '10k',
      is_test_page_set: false,
      repeat_runs: 1,
      run_in_parallel: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --max-pages-per-bot=10',
      browser_args_no_patch: '',
      browser_args_with_patch: '',
      description:
        'Testing https://chromium-review.googlesource.com/c/2316326/1 ([WIP] Track dynamically added scripts with no src in AdTracker)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/8e2c009191f9ac50e59b8c896f475e2a4c555b59.patch',
      blink_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_base_build_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      results:
        'https://ct.skia.org/results/cluster-telemetry/tasks/chromium_perf_runs/alexmt-ChromiumPerf-5091/html/index.html',
      no_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5091-nopatch/consolidated_outputs/alexmt-ChromiumPerf-5091-nopatch.output',
      with_patch_raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumPerf-5091-withpatch/consolidated_outputs/alexmt-ChromiumPerf-5091-withpatch.output',
      chromium_hash: '',
      cc_list: null,
      task_priority: 100,
      group_name: '',
      value_column_name: 'sum',
    },
  ],
  ids: [5094, 5091, 5084, 5082, 5081, 5074, 5065, 5063, 5062, 5061],
  pagination: { offset: 20, size: 10, total: 1827 },
  permissions: [
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
  ],
};
