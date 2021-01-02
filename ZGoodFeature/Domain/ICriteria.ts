export interface IFilter {
  readonly filterField: string;
  readonly filterOperator: string;
  readonly filterValue: string;
}

export interface ICriteria {
  filters: IFilter[];
  order: string;
  offset: number;
  limit: number;
}
