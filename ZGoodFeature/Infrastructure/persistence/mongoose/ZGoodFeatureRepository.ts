import { ZGoodFeatureModel } from ".";
import { ZGoodFeature, IZGoodFeatureRepository, ICriteria, IZGoodFeature } from "../../../Domain";

export class ZGoodFeatureRepository implements IZGoodFeatureRepository {
	constructor() {}

	async query(criteria: ICriteria): Promise<{ data: IZGoodFeature[]; count: number }> {
		const [data, count] = await Promise.all([
			ZGoodFeatureModel.find(criteria.serializeFilter())
				.skip(criteria.offset)
				.limit(criteria.limit),
			ZGoodFeatureModel.countDocuments(criteria.serializeFilter())
		]);

		return { data, count };
	}

	async add(zGoodFeature: ZGoodFeature): Promise<void> {
		await ZGoodFeatureModel.create(zGoodFeature);
		return;
	}

	async update(zGoodFeature: ZGoodFeature): Promise<void> {
		const { id } = zGoodFeature;
		delete zGoodFeature.id;

		await ZGoodFeatureModel.findByIdAndUpdate(id, zGoodFeature);
		return;
	}

	async remove(id: string): Promise<void> {
		await ZGoodFeatureModel.findByIdAndDelete(id);
		return;
	}
}
