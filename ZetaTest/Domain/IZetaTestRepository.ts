import { ICriteria, IZetaTest, ZetaTest } from '.';

export interface IZetaTestRepository {
	add(zetaTest: ZetaTest): Promise<void>;
	update(zetaTest: ZetaTest): Promise<void>;
	remove(id: string): Promise<void>;
	query(criteria: ICriteria): Promise<{ data: IZetaTest[]; count: number }>;
}
