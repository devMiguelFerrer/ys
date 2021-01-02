package infrastructure

import "fmt"

func handleTypeORMModel() string {
	return `import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';
import { I__name__ } from '../../../Domain';

@Entity('__var__')
export class __name__Model implements I__name__ {
`
}

func handleTypeORMCriteria() string {
	return `import { IFilter, ICriteria } from '../../../Domain';

export class __name__Criteria implements ICriteria {
	readonly offset: number;
	readonly limit: number;
	readonly filters: IFilter[];
	readonly order: string;

	constructor(query: { filter: string }) {
			const decodedFilters = this.decodedFilter(query.filter);
			this.offset = decodedFilters.offset && !Number.isNaN(parseInt(decodedFilters.offset, 10)) ? parseInt(decodedFilters.offset, 10) : 0;
			this.limit = decodedFilters.limit && !Number.isNaN(parseInt(decodedFilters.limit, 10)) ? parseInt(decodedFilters.limit, 10) : 0;
			this.filters = decodedFilters.filters && Array.isArray(decodedFilters.filters) ? [...decodedFilters.filters] : [];
	}

	private decodedFilter(encodedFilters: string): { offset: string; limit: string; filters: [] } {
		try {
				return JSON.parse(decodeURI(encodedFilters));
		} catch (error) {
				return { offset: '0', limit: '0', filters: [] };
		}
	}

}
`
}

func handleTypeORMRepository() string {
	repositoryModel := `import { createQueryBuilder, getConnection, SelectQueryBuilder } from 'typeorm';
import { __name__Model } from '.';
import { I__name__, I__name__Repository, ICriteria } from '../../../Domain';

export class __name__Repository implements I__name__Repository {
	constructor() {}

	async add(__var__: I__name__): Promise<void> {
		await getConnection().getRepository(__name__Model).save(__var__);
		return;
	}

	async update(updated__name__: I__name__): Promise<void> {
		let old__name__ = await getConnection().getRepository(__name__Model).findOne(updated__name__.id);
		old__name__ = { ...old__name__, ...updated__name__ };
		await getConnection().getRepository(__name__Model).save(old__name__);
		return;
	}

	async remove(id: string): Promise<void> {
		await getConnection().getRepository(__name__Model).delete(id);
		return;
	}

	async query(criteria: ICriteria): Promise<{ data: I__name__[]; count: number }> {
		const queryBuilder = getConnection()
			.getRepository(__name__Model)
			.createQueryBuilder('__name__');
		
		const queryFiltered = this.handleCriteria(queryBuilder, criteria);
		
		const [data, count] = await Promise.all([
			queryFiltered.getMany(),
			queryFiltered.getCount(),
		]);

		return { data, count };
	}

	private handleCriteria(query: SelectQueryBuilder<__name__Model>, criteria: ICriteria): SelectQueryBuilder<__name__Model> {
		criteria.filters.forEach((filter) => {
			if (filter.filterOperator === 'OR') {`
	repositoryModel += fmt.Sprintln("query.orWhere(`__name__.${filter.filterField} LIKE :value`, { value: `%${filter.filterValue}%` })")
	repositoryModel += `			}
		})
		query.limit(criteria.limit).offset(criteria.offset);

		return query;
	}
}	
`
	return repositoryModel
}
