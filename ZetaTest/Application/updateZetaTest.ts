import { IZetaTestRepository, IZetaTest } from '../Domain';

export class UpdateZetaTest {
	constructor(private repository: IZetaTestRepository, private zetaTest: IZetaTest) {}

		async dispatch(): Promise<void> {
			await this.repository.update(this.zetaTest);
			return;
		}
}
