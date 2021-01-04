import { IZetaTestRepository, IZetaTest } from '../Domain';

export class AddZetaTest {
	constructor(private repository: IZetaTestRepository, private zetaTest: IZetaTest) {}

		async dispatch(): Promise<void> {
			await this.repository.add(this.zetaTest);
			return;
		}
}
