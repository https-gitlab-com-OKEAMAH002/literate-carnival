import { GetTasksResponse } from '../json';

export const tasksResult0: GetTasksResponse = {
  data: [
    {
      ts_added: 20200728124631,
      ts_started: 20200728124633,
      ts_completed: 20200728163231,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 1,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2425\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4dae475910ce3d10',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args:
        '--force-fieldtrials=Foo/Enabled --force-fieldtrial-params=Foo.Enabled:activation_state/enabled/activation_scope/all_sites/activation_priority/10000 --enable-features=SubresourceFilter\u003cFoo',
      description: 'Regular AdTagging accuracy run with blocking on',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2425/consolidated_outputs/alexmt-ChromiumAnalysis-2425.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTaggingWithBlocking',
    },
    {
      ts_added: 20200728124131,
      ts_started: 20200728124133,
      ts_completed: 20200728152231,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 1,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2424\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4dae42c5d3a33210',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args: '',
      description: 'Regular AdTagging accuracy run',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2424/consolidated_outputs/alexmt-ChromiumAnalysis-2424.output',
      value_column_name: 'sum',
      MatchStdoutTxt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTagging',
    },
    {
      ts_added: 20200727124903,
      ts_started: 20200727124905,
      ts_completed: 20200727163048,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 2,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:dproy-ChromiumAnalysis-2423\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da9234d87570110',
      benchmark: 'loading.cluster_telemetry',
      page_sets: 'VoltMobile10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --traffic-setting=Regular-4G --use-live-sites',
      browser_args: '',
      description: 'Regular run for Volt 10k pages (Chrome M80, live sites)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Android',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/dproy-ChromiumAnalysis-2423/consolidated_outputs/dproy-ChromiumAnalysis-2423.output',
      value_column_name: 'avg',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath:
        'gs://chrome-unsigned/android-B0urB0N/80.0.3987.87/arm_64/Monochrome.apk',
      telemetry_isolate_hash: '48afa0e9204793e1d3d4615e1a7ba26174de77d4',
      cc_list: ['someone@google.com'],
      task_priority: 100,
      group_name: 'volt10k-m80',
    },
    {
      ts_added: 20200727124902,
      ts_started: 20200727124904,
      ts_completed: 20200727163448,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2422\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da92349b67fc510',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args:
        '--force-fieldtrials=Foo/Enabled --force-fieldtrial-params=Foo.Enabled:activation_state/enabled/activation_scope/all_sites/activation_priority/10000 --enable-features=SubresourceFilter\u003cFoo',
      description: 'Regular AdTagging accuracy run with blocking on',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2422/consolidated_outputs/alexmt-ChromiumAnalysis-2422.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTaggingWithBlocking',
    },
    {
      ts_added: 20200727124402,
      ts_started: 20200727124406,
      ts_completed: 20200727154048,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2421\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da91ebe58bd0c10',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args: '',
      description: 'Regular AdTagging accuracy run',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2421/consolidated_outputs/alexmt-ChromiumAnalysis-2421.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTagging',
    },
    {
      ts_added: 20200726125152,
      ts_started: 20200726125154,
      ts_completed: 20200726145552,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2420\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da3ff85b89e5610',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args:
        '--force-fieldtrials=Foo/Enabled --force-fieldtrial-params=Foo.Enabled:activation_state/enabled/activation_scope/all_sites/activation_priority/10000 --enable-features=SubresourceFilter\u003cFoo',
      description: 'Regular AdTagging accuracy run with blocking on',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2420/consolidated_outputs/alexmt-ChromiumAnalysis-2420.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTaggingWithBlocking',
    },
    {
      ts_added: 20200726124652,
      ts_started: 20200726124655,
      ts_completed: 20200726173152,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2419\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da3faf58ae89d10',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args: '',
      description: 'Regular AdTagging accuracy run',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2419/consolidated_outputs/alexmt-ChromiumAnalysis-2419.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTagging',
    },
    {
      ts_added: 20200725125652,
      ts_started: 20200725125653,
      ts_completed: 20200725165152,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2418\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d9eddbd1aaf8210',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args:
        '--force-fieldtrials=Foo/Enabled --force-fieldtrial-params=Foo.Enabled:activation_state/enabled/activation_scope/all_sites/activation_priority/10000 --enable-features=SubresourceFilter\u003cFoo',
      description: 'Regular AdTagging accuracy run with blocking on',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2418/consolidated_outputs/alexmt-ChromiumAnalysis-2418.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTaggingWithBlocking',
    },
    {
      ts_added: 20200725125153,
      ts_started: 20200725125155,
      ts_completed: 20200725163352,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:dproy-ChromiumAnalysis-2417\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d9ed92d95aea710',
      benchmark: 'loading.cluster_telemetry',
      page_sets: 'VoltMobile10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --traffic-setting=Regular-4G --use-live-sites',
      browser_args: '',
      description: 'Regular run for Volt 10k pages (Chrome M80, live sites)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Android',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/dproy-ChromiumAnalysis-2417/consolidated_outputs/dproy-ChromiumAnalysis-2417.output',
      value_column_name: 'avg',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath:
        'gs://chrome-unsigned/android-B0urB0N/80.0.3987.87/arm_64/Monochrome.apk',
      telemetry_isolate_hash: '48afa0e9204793e1d3d4615e1a7ba26174de77d4',
      cc_list: ['someone@google.com'],
      task_priority: 100,
      group_name: 'volt10k-m80',
    },
    {
      ts_added: 20200725125152,
      ts_started: 20200725125154,
      ts_completed: 20200725154952,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2416\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d9ed92b6fc90410',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args: '',
      description: 'Regular AdTagging accuracy run',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2416/consolidated_outputs/alexmt-ChromiumAnalysis-2416.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTagging',
    },
  ],
  ids: [2425, 2424, 2423, 2422, 2421, 2420, 2419, 2418, 2417, 2416],
  pagination: { offset: 0, size: 10, total: 1766 },
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
      ts_added: 20200726124652,
      ts_started: 20200726124655,
      ts_completed: 20200726173152,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2419\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da3faf58ae89d10',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args: '',
      description: 'Regular AdTagging accuracy run',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2419/consolidated_outputs/alexmt-ChromiumAnalysis-2419.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTagging',
    },
    {
      ts_added: 20200725125652,
      ts_started: 20200725125653,
      ts_completed: 20200725165152,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2418\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d9eddbd1aaf8210',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args:
        '--force-fieldtrials=Foo/Enabled --force-fieldtrial-params=Foo.Enabled:activation_state/enabled/activation_scope/all_sites/activation_priority/10000 --enable-features=SubresourceFilter\u003cFoo',
      description: 'Regular AdTagging accuracy run with blocking on',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2418/consolidated_outputs/alexmt-ChromiumAnalysis-2418.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTaggingWithBlocking',
    },
    {
      ts_added: 20200725125153,
      ts_started: 20200725125155,
      ts_completed: 20200725163352,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:dproy-ChromiumAnalysis-2417\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d9ed92d95aea710',
      benchmark: 'loading.cluster_telemetry',
      page_sets: 'VoltMobile10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --pageset-repeat=1 --skip-typ-expectations-tags-validation --legacy-json-trace-format --traffic-setting=Regular-4G --use-live-sites',
      browser_args: '',
      description: 'Regular run for Volt 10k pages (Chrome M80, live sites)',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Android',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/dproy-ChromiumAnalysis-2417/consolidated_outputs/dproy-ChromiumAnalysis-2417.output',
      value_column_name: 'avg',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath:
        'gs://chrome-unsigned/android-B0urB0N/80.0.3987.87/arm_64/Monochrome.apk',
      telemetry_isolate_hash: '48afa0e9204793e1d3d4615e1a7ba26174de77d4',
      cc_list: ['someone@google.com'],
      task_priority: 100,
      group_name: 'volt10k-m80',
    },
    {
      ts_added: 20200725125152,
      ts_started: 20200725125154,
      ts_completed: 20200725154952,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2416\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4d9ed92b6fc90410',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args: '',
      description: 'Regular AdTagging accuracy run',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2416/consolidated_outputs/alexmt-ChromiumAnalysis-2416.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTagging',
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
      ts_added: 20200726125152,
      ts_started: 20200726125154,
      ts_completed: 20200726145552,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2420\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da3ff85b89e5610',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args:
        '--force-fieldtrials=Foo/Enabled --force-fieldtrial-params=Foo.Enabled:activation_state/enabled/activation_scope/all_sites/activation_priority/10000 --enable-features=SubresourceFilter\u003cFoo',
      description: 'Regular AdTagging accuracy run with blocking on',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2420/consolidated_outputs/alexmt-ChromiumAnalysis-2420.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTaggingWithBlocking',
    },
    {
      ts_added: 20200726124652,
      ts_started: 20200726124655,
      ts_completed: 20200726173152,
      username: 'someone@google.com',
      failure: false,
      repeat_after_days: 0,
      swarming_logs:
        'https://chrome-swarming.appspot.com/tasklist?l=500\u0026c=name\u0026c=created_ts\u0026c=bot\u0026c=duration\u0026c=state\u0026f=runid:alexmt-ChromiumAnalysis-2419\u0026st=1262304000000',
      task_done: true,
      swarming_task_id: '4da3faf58ae89d10',
      benchmark: 'ad_tagging.cluster_telemetry',
      page_sets: '10k',
      is_test_page_set: false,
      benchmark_args:
        '--output-format=csv --skip-typ-expectations-tags-validation --legacy-json-trace-format --verbose-cpu-metrics --verbose-memory-metrics',
      browser_args: '',
      description: 'Regular AdTagging accuracy run',
      custom_webpages_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      chromium_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      skia_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      catapult_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      benchmark_patch_gspath:
        'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      v8_patch_gspath: 'patches/da39a3ee5e6b4b0d3255bfef95601890afd80709.patch',
      run_in_parallel: false,
      platform: 'Linux',
      run_on_gce: false,
      raw_output:
        'https://ct.skia.org/results/cluster-telemetry/tasks/benchmark_runs/alexmt-ChromiumAnalysis-2419/consolidated_outputs/alexmt-ChromiumAnalysis-2419.output',
      value_column_name: 'sum',
      match_stdout_txt: '',
      chromium_hash: '',
      apk_gspath: '',
      telemetry_isolate_hash: '',
      cc_list: null,
      task_priority: 110,
      group_name: 'AdTagging',
    },
  ],
  ids: [5065, 5063],
  pagination: { offset: 20, size: 10, total: 1827 },
  permissions: [
    { DeleteAllowed: true, RedoAllowed: true },
    { DeleteAllowed: true, RedoAllowed: true },
  ],
};
