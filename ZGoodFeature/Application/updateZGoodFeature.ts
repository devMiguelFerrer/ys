import { IZGoodFeatureRepository, IZGoodFeature } from '../Domain';

export class UpdateZGoodFeature {
	constructor(private repository: IZGoodFeatureRepository, private zGoodFeature: IZGoodFeature) {}

		async dispatch(): Promise<void> {
			await this.repository.update(this.zGoodFeature);
			return;
		}
}
