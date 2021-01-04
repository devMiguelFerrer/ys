import { ValidationBoolean, ValidationNumber, ValidationString } from './BaseTypes';

export class Age extends ValidationNumber {
	constructor(value: number) {
		super(value);
		if (this.min(0) || this.max(999999)) {
			throw new Error('The value should be a number');
		}
  }
}

export class Active extends ValidationBoolean {
	constructor(value: boolean) {
		super(value);
  }
}

export class Firstname extends ValidationString {
	constructor(value: string) {
		super(value);
		if (this.min(1) || this.max(50)) {
			throw new Error('The value should be a number');
		}
  }
}

export class Lastname extends ValidationString {
	constructor(value: string) {
		super(value);
		if (this.min(1) || this.max(50)) {
			throw new Error('The value should be a number');
		}
  }
}
