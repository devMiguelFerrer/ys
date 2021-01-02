import { ICriteria, IFilter } from "../../../Domain";

enum OperatorMongoose {
	OR = "$or",
	AND = "$and",
	IN = "$in"
}

export class ZGoodFeatureCriteria implements ICriteria {
	readonly offset: number;

	readonly limit: number;

	readonly filters: IFilter[];

	readonly order: string;

	constructor(query: { filter: string }) {
		const decodedFilters = decodedFilter(query.filter);
		this.offset = decodedFilters.offset && !Number.isNaN(parseInt(decodedFilters.offset, 10)) ? parseInt(decodedFilters.offset, 10) : 0;
		this.limit = decodedFilters.limit && !Number.isNaN(parseInt(decodedFilters.limit, 10)) ? parseInt(decodedFilters.limit, 10) : 0;
		this.filters = decodedFilters.filters && Array.isArray(decodedFilters.filters) ? [...decodedFilters.filters] : [];
	}

	public serializeFilter() {
		const serializeFilter = {};
		let flag = "";
		this.filters.map(filter => {
			const { filterField, filterOperator, filterValue } = filter;
			if (!serializeFilter[OperatorMongoose[filterOperator]]) {
				serializeFilter[OperatorMongoose[filterOperator]] = [];
			}
			if (typeof filterValue === "string") {
				flag = "i";
			}
			serializeFilter[OperatorMongoose[filterOperator]].push({ [filterField]: new RegExp(filterValue, flag) });
		});
		return serializeFilter;
	}
}

function decodedFilter(encodedFilters: string): { offset: string; limit: string; filters: [] } {
	try {
		return JSON.parse(decodeURI(encodedFilters));
	} catch (error) {
		return { offset: "0", limit: "0", filters: [] };
	}
}
