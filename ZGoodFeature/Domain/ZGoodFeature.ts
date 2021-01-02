import { IZGoodFeature } from '.';

export class ZGoodFeature implements IZGoodFeature {
  firstname: string
  lastname: string
  age: number
  active: boolean
  constructor(zGoodFeature: IZGoodFeature) {
    this.firstname = zGoodFeature.firstname
    this.lastname = zGoodFeature.lastname
    this.age = zGoodFeature.age
    this.active = zGoodFeature.active
  }
}
