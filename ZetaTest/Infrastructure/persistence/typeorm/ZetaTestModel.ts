import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';
import { IZetaTest } from '../../../Domain';

@Entity('zetaTest')
export class ZetaTestModel implements IZetaTest {
  @Column()
  firstname: string

  @Column()
  lastname: string

  @Column()
  age: number

  @Column()
  active: boolean

  constructor(zetaTest: IZetaTest) {
    this.firstname = zetaTest.firstname
    this.lastname = zetaTest.lastname
    this.age = zetaTest.age
    this.active = zetaTest.active
  }
}
