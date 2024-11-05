// DO NOT EDIT. This file is automatically generated.

export interface Changelist {
	system: string;
	id: string;
	owner: string;
	status: string;
	subject: string;
	updated: string;
	url: string;
}

export interface TryJob {
	id: string;
	name: string;
	updated: string;
	system: string;
	url: string;
}

export interface Patchset {
	id: string;
	order: number;
	try_jobs: TryJob[];
}

export interface ChangelistSummaryResponse {
	cl: Changelist;
	patch_sets: Patchset[];
	num_total_patch_sets: number;
}

export interface TriageHistory {
	user: string;
	ts: string;
}

export interface Trace {
	label: TraceID;
	data: number[] | null;
	params: { [key: string]: string } | null;
	comment_indices: number[] | null;
}

export interface DigestStatus {
	digest: Digest;
	status: Label;
}

export interface TraceGroup {
	traces: Trace[] | null;
	digests: DigestStatus[] | null;
	total_digests: number;
}

export interface SRDiffDigest {
	numDiffPixels: number;
	combinedMetric: number;
	pixelDiffPercent: number;
	maxRGBADiffs: number[];
	dimDiffer: boolean;
	digest: Digest;
	status: Label;
	paramset: ParamSet;
}

export interface SearchResult {
	digest: Digest;
	test: TestName;
	status: Label;
	triage_history: TriageHistory[] | null;
	paramset: ParamSet;
	traces: TraceGroup;
	refDiffs: { [key: string]: SRDiffDigest | null } | null;
	closestRef: RefClosest;
}

export interface Commit {
	commit_time: number;
	id: string;
	hash: string;
	author: string;
	message: string;
	cl_url: string;
}

export interface BulkTriageDeltaInfo {
	grouping: Params;
	digest: Digest;
	label_before: Label;
	closest_diff_label: ClosestDiffLabel;
	in_current_search_results_page: boolean;
}

export interface SearchResponse {
	digests: (SearchResult | null)[] | null;
	offset: number;
	size: number;
	commits: Commit[] | null;
	bulk_triage_delta_infos: BulkTriageDeltaInfo[];
}

export interface TriageRequest {
	testDigestStatus: TriageRequestData;
	changelist_id: string;
	crs: string;
	imageMatchingAlgorithm?: string;
}

export interface TriageDelta {
	grouping: Params;
	digest: Digest;
	label_before: Label;
	label_after: Label;
}

export interface TriageRequestV3 {
	deltas: TriageDelta[];
	changelist_id?: string;
	crs?: string;
	image_matching_algorithm?: string;
}

export interface TriageConflict {
	grouping: Params;
	digest: Digest;
	expected_label_before: Label;
	actual_label_before: Label;
}

export interface TriageResponse {
	status: TriageResponseStatus;
	conflict?: TriageConflict;
}

export interface GUICorpusStatus {
	name: string;
	untriagedCount: number;
}

export interface StatusResponse {
	lastCommit: Commit;
	corpStatus: GUICorpusStatus[];
}

export interface GroupingsResponse {
	grouping_param_keys_by_corpus: { [key: string]: string[] | null } | null;
}

export interface TestRollup {
	grouping: Params;
	num: number;
	sample_digest: Digest;
}

export interface ByBlameEntry {
	groupID: string;
	nDigests: number;
	nTests: number;
	affectedTests: TestRollup[] | null;
	commits: Commit[] | null;
}

export interface ByBlameResponse {
	data: ByBlameEntry[] | null;
}

export interface TriageLogEntry {
	id: string;
	name: string;
	ts: number;
	details: TriageDelta[];
}

export interface TriageLogResponse {
	entries: TriageLogEntry[];
	offset: number;
	size: number;
	total: number;
}

export interface ChangelistsResponse {
	changelists: Changelist[] | null;
	offset: number;
	size: number;
	total: number;
}

export interface IgnoreRuleBody {
	duration: string;
	filter: string;
	note: string;
}

export interface IgnoreRule {
	id: string;
	name: string;
	updatedBy: string;
	expires: string;
	query: string;
	note: string;
	countAll: number;
	exclusiveCountAll: number;
	count: number;
	exclusiveCount: number;
}

export interface IgnoresResponse {
	rules: IgnoreRule[] | null;
}

export interface TestSummary {
	grouping: Params;
	positive_digests: number;
	negative_digests: number;
	untriaged_digests: number;
	total_digests: number;
}

export interface ListTestsResponse {
	tests: TestSummary[] | null;
}

export interface LeftDiffInfo {
	test: TestName;
	digest: Digest;
	status: Label;
	triage_history: TriageHistory[] | null;
	paramset: ParamSet;
}

export interface DigestComparison {
	left: LeftDiffInfo;
	right: SRDiffDigest;
}

export interface DetailsRequest {
	grouping: Params;
	digest: Digest;
	changelist_id?: string;
	crs?: string;
}

export interface DigestDetails {
	digest: SearchResult;
	commits: Commit[] | null;
}

export interface ClusterDiffNode {
	name: Digest;
	status: Label;
}

export interface ClusterDiffLink {
	source: number;
	target: number;
	value: number;
}

export interface ClusterDiffResult {
	nodes: ClusterDiffNode[] | null;
	links: ClusterDiffLink[] | null;
	test: TestName;
	paramsetByDigest: { [key: string]: ParamSet };
	paramsetsUnion: ParamSet;
}

export interface DiffRequest {
	grouping: Params;
	left_digest: Digest;
	right_digest: Digest;
	changelist_id?: string;
	crs?: string;
}

export interface GroupingForTestRequest {
	test_name: string;
}

export interface GroupingForTestResponse {
	grouping: Params;
}

export type ParamSet = { [key: string]: string[] };

export type ParamSetResponse = { [key: string]: string[] | null } | null;

export type Digest = string;

export type TestName = string;

export type Label = 'untriaged' | 'positive' | 'negative';

export type TraceID = string;

export type RefClosest = 'pos' | 'neg' | '';

export type TriageRequestData = { [key: string]: { [key: string]: Label } | null } | null;

export type Params = { [key: string]: string };

export type ClosestDiffLabel = 'none' | 'untriaged' | 'positive' | 'negative';

export type TriageResponseStatus = 'ok' | 'conflict';
