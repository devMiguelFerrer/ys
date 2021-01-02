import { IZGoodFeatureRepository } from '../Domain';

export class RemoveZGoodFeature {
	constructor(private repository: IZGoodFeatureRepository, private id: string) {}

		async dispatch(): Promise<void> {
			await this.repository.remove(this.id);
			return;
		}
}
