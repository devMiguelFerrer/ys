import { IZGoodFeatureRepository, ICriteria, IZGoodFeature } from '../Domain';

export class GetZGoodFeature {
		constructor(private repository: IZGoodFeatureRepository, private criteria: ICriteria) {}

		async dispatch(): Promise<{ data: IZGoodFeature[]; count: number }> {
				return this.repository.query(this.criteria);
		}
}
