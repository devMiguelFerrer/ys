import { IZetaTestRepository } from '../Domain';

export class RemoveZetaTest {
	constructor(private repository: IZetaTestRepository, private id: string) {}

		async dispatch(): Promise<void> {
			await this.repository.remove(this.id);
			return;
		}
}
