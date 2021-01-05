package domain

func handleValidationString() string {
	return `export class ValidationString {
	value: string;

	constructor(value: string) {
		if (typeof value !== 'string') {
			throw new Error('The value should be a string');
		}
		this.value = value;
	}

	min(minValue: number): boolean {
			return this.value.trim().length < minValue;
	}

	max(minValue: number): boolean {
			return this.value.trim().length > minValue;
	}
}
`
}

func handleValidationNumber() string {
	return `export class ValidationNumber {
	value: number;

	constructor(value: number) {
		if (typeof value !== 'number') {
			throw new Error('The value should be a number');
		}
		this.value = value;
	}

	min(minValue: number): boolean {
			return this.value < minValue;
	}

	max(minValue: number): boolean {
			return this.value > minValue;
	}
}
`
}

func handleValidationBoolean() string {
	return `export class ValidationBoolean {
	value: boolean;

	constructor(value: boolean) {
		if (typeof value !== 'boolean') {
			throw new Error('The value should be a boolean');
		}
		this.value = value;
	}
}
`
}
