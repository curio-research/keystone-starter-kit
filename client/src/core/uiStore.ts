// Stores helper UI states

const defaultStartingDirection = 'up';

export class UIState {
  public selectedTableToDisplay: string;
  public lastMovedDirection: string;

  constructor() {
    this.selectedTableToDisplay = '';
    this.lastMovedDirection = defaultStartingDirection;
  }

  public setSelectedTableToDisplay(tableName: string) {
    this.selectedTableToDisplay = tableName;
  }
}
