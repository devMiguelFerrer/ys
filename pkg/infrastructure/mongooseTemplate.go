package infrastructure

func handleMongooseModel() string {
	// TODO:
	return ``
}

func handleMongooseCriteria() string {
	// TODO:
	return ``
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
