// ui store
export class UIStore {
  public selectedTableToDisplay: string;
  public lastMovedDirection: string;

  constructor() {
    this.selectedTableToDisplay = '';
    this.lastMovedDirection = '';
  }

  public setSelectedTableToDisplay(tableName: string) {
    this.selectedTableToDisplay = tableName;
  }
}
