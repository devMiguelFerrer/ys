import { createQueryBuilder, getConnection, SelectQueryBuilder } from 'typeorm';
import { ZetaTestModel } from '.';
import { IZetaTest, IZetaTestRepository, ICriteria } from '../../../Domain';

export class ZetaTestRepository implements IZetaTestRepository {
	constructor() {}

	async add(zetaTest: IZetaTest): Promise<void> {
		await getConnection().getRepository(ZetaTestModel).save(zetaTest);
		return;
	}

	async update(updatedZetaTest: IZetaTest): Promise<void> {
		let oldZetaTest = await getConnection().getRepository(ZetaTestModel).findOne(updatedZetaTest.id);
		oldZetaTest = { ...oldZetaTest, ...updatedZetaTest };
		await getConnection().getRepository(ZetaTestModel).save(oldZetaTest);
		return;
	}

	async remove(id: string): Promise<void> {
		await getConnection().getRepository(ZetaTestModel).delete(id);
		return;
	}

	async query(criteria: ICriteria): Promise<{ data: IZetaTest[]; count: number }> {
		const queryBuilder = getConnection()
			.getRepository(ZetaTestModel)
			.createQueryBuilder('ZetaTest');
		
		const queryFiltered = this.handleCriteria(queryBuilder, criteria);
		
		const [data, count] = await Promise.all([
			queryFiltered.getMany(),
			queryFiltered.getCount(),
		]);

		return { data, count };
	}

	private handleCriteria(query: SelectQueryBuilder<ZetaTestModel>, criteria: ICriteria): SelectQueryBuilder<ZetaTestModel> {
		criteria.filters.forEach((filter) => {
			if (filter.filterOperator === 'OR') {query.orWhere(`ZetaTest.${filter.filterField} LIKE :value`, { value: `%${filter.filterValue}%` })
			}
		})
		query.limit(criteria.limit).offset(criteria.offset);

		return query;
	}
}	
