// ui store
export class UIStore {
  public selectedTableToDisplay: string;

  constructor() {
    this.selectedTableToDisplay = "";
  }

  public setSelectedTableToDisplay(tableName: string) {
    this.selectedTableToDisplay = tableName;
  }
}
