import { ICriteria, IZGoodFeature, ZGoodFeature } from ".";

  export interface IZGoodFeatureRepository {
    add(zGoodFeature: ZGoodFeature): Promise<void>;
    update(zGoodFeature: ZGoodFeature): Promise<void>;
    remove(id: string): Promise<void>;
    query(criteria: ICriteria): Promise<{ data: IZGoodFeature[]; count: number }>;
  }
