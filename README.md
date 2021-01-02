# YES, SIR!!! 🎩 (ys)
CLI to build templates for Typescript.

### Commands
```
ys create [ResourceName] -e [EntityFile] -r [ModelType]
```
```
ys create CustomTask -e config.json -r mongoose
```
### Folder Structure
```
📦CustomTask
  ┣ 🗂️Application
  ┃ ┣ 📃addCustomTask.ts
  ┃ ┣ 📃getCustomTask.ts
  ┃ ┣ 📃index.ts
  ┃ ┣ 📃removeCustomTask.ts
  ┃ ┗ 📃updateCustomTask.ts
  ┣ 🗂️Domain
  ┃ ┣ 📃ICriteria.ts
  ┃ ┣ 📃ICustomTask.ts
  ┃ ┣ 📃ICustomTaskRepository.ts
  ┃ ┣ 📃index.ts
  ┃ ┗ 📃CustomTask.ts
  ┗ 🗂️Infrastructure
    ┣ 📂Controllers
    ┃ ┗ 📃customTaskController.ts
    ┗ 📂Persistence
      ┗ 📂[typeorm || mongoose]
        ┣ 📃index.ts
        ┣ 📃customTaskCriteria.ts
        ┣ 📃customTaskModel.ts
        ┗ 📃customTaskRepository.ts
```
