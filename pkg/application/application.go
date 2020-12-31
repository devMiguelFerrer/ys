package application

import "github.com/devMiguelFerrer/ys/pkg/shared"

// GenerateAdd method create ADD Use Case
func GenerateAdd(recipe []shared.Recipe) string {
	const base = `import { I__name__Repository, I__name__ } from '../Domain';

export class Add__name__ {
	constructor(private repository: I__name__Repository, private __var__: I__name__) {}

		async dispatch(): Promise<void> {
			await this.repository.add(this.__var__);
			return;
		}
}
`
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateDelete method create REMOVE Use Case
func GenerateDelete(recipe []shared.Recipe) string {
	const base = `import { I__name__Repository } from '../Domain';

export class Remove__name__ {
	constructor(private repository: I__name__Repository, private id: string) {}

		async dispatch(): Promise<void> {
			await this.repository.remove(this.id);
			return;
		}
}
`
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateGet method create GET Use Case
func GenerateGet(recipe []shared.Recipe) string {
	const base = `import { I__name__Repository, ICriteria, I__name__ } from '../Domain';

export class Get__name__ {
		constructor(private repository: I__name__Repository, private criteria: ICriteria) {}

		async dispatch(): Promise<{ data: I__name__[]; count: number }> {
				return this.repository.query(this.criteria);
		}
}
`

	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateUpdate method create UPDATE Use Case
func GenerateUpdate(recipe []shared.Recipe) string {
	const base = `import { I__name__Repository, I__name__ } from '../Domain';

export class Update__name__ {
	constructor(private repository: I__name__Repository, private __var__: I__name__) {}

		async dispatch(): Promise<void> {
			await this.repository.update(this.__var__);
			return;
		}
}
`
	return shared.ReplaceByRecipe(base, recipe)
}

// GenerateApplicationIndex method create INDEX to export Use Cases
func GenerateApplicationIndex(recipe []shared.Recipe) string {
	const base = `export * from './add__name__';
export * from './get__name__';
export * from './remove__name__';
export * from './update__name__';
`
	return shared.ReplaceByRecipe(base, recipe)
}
