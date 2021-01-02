import { IZGoodFeatureRepository, IZGoodFeature } from '../Domain';

export class AddZGoodFeature {
	constructor(private repository: IZGoodFeatureRepository, private zGoodFeature: IZGoodFeature) {}

		async dispatch(): Promise<void> {
			await this.repository.add(this.zGoodFeature);
			return;
		}
}
