package domain

func handleValidationString() string {
	return `class ValidationString {
	value: string;
	constructor(value: string) {
		if (typeof value !== 'string') {
			throw new Error('The should be a string');
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
	return `class ValidationNumber {
	value: number;
	constructor(value: number) {
		if (typeof value !== 'number') {
			throw new Error('The should be a number');
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
