import { IZetaTestRepository, ICriteria, IZetaTest } from '../Domain';

export class GetZetaTest {
		constructor(private repository: IZetaTestRepository, private criteria: ICriteria) {}

		async dispatch(): Promise<{ data: IZetaTest[]; count: number }> {
				return this.repository.query(this.criteria);
		}
}
