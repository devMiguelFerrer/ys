package infrastructure

func handleMongooseModel(columns string) string {
	repositoryModel := `import * as mongoose from "mongoose";
import { I__name__ } from "../../../Domain";

interface __name__Doc extends mongoose.Document, I__name__ {
	id: string;
	username: string;
	group_id: string;
}

type __name__Model = mongoose.Model<__name__Doc>

const __var__Schema = new mongoose.Schema<I__name__>(
	{
		id_user: {
			type: String,
			required: true
		},
		username: {
			type: String,
			required: true
		},
		group_id: {
			type: String,
			required: true
		},
`
	repositoryModel += columns
	repositoryModel += `  },
	{
		toJSON: {
			transform(doc, ret) {
				delete ret.__v;
				ret.id = ret._id;
			}
		}
	}
);

const __name__Model = mongoose.model<__name__Doc, __name__Model>("__var__", __var__Schema);

export { __name__Model };
`
	return repositoryModel
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
