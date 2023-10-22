// ui store
export class UIState {
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
