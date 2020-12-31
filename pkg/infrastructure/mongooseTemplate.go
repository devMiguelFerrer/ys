package infrastructure

func handleMongooseModel() string {
	// TODO:
	return ``
}

func handleMongooseCriteria() string {
	// TODO:
	return `import { ICriteria, IFilter } from "../../../Domain";

enum OperatorMongoose {
	OR = "$or",
	AND = "$and",
	IN = "$in"
}

export class __name__Criteria implements ICriteria {
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
`
}

func handleMongooseRepository() string {
	return `import { __name__Model } from ".";
import { __name__, I__name__Repository, ICriteria, I__name__ } from "../../../Domain";

export class __name__Repository implements I__name__Repository {
	constructor() {}

	async query(criteria: ICriteria): Promise<{ data: I__name__[]; count: number }> {
		const [data, count] = await Promise.all([
			__name__Model.find(criteria.serializeFilter())
				.skip(criteria.offset)
				.limit(criteria.limit),
			__name__Model.countDocuments(criteria.serializeFilter())
		]);

		return { data, count };
	}

	async add(__var__: __name__): Promise<void> {
		await __name__Model.create(__var__);
		return;
	}

	async update(__var__: __name__): Promise<void> {
		const { id } = __var__;
		delete __var__.id;

		await __name__Model.findByIdAndUpdate(id, __var__);
		return;
	}

	async remove(id: string): Promise<void> {
		await __name__Model.findByIdAndDelete(id);
		return;
	}
}
`
}
