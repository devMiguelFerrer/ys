import { IZetaTest } from '.';

export class ZetaTest implements IZetaTest {
  firstname: string
  lastname: string
  age: number
  active: boolean
  constructor(zetaTest: IZetaTest) {
    this.firstname = new Firstname(zetaTest.firstname).value
    this.lastname = new Lastname(zetaTest.lastname).value
    this.age = new Age(zetaTest.age).value
    this.active = new Active(zetaTest.active).value
  }
}
