import { Position } from 'core/schemas';

export class PositionWrapperManager {
  private entitiesToLocalPositions = new Map<number, Position>();
  private entitiesToTargetPositions = new Map<number, Position>();

  constructor() {
    setInterval(() => {
      this.updateAll();
    }, 10);
  }

  public setTargetPosition(entity: number, position: Position) {
    this.entitiesToTargetPositions.set(entity, position);
  }

  private setLocalPosition(entity: number, position: Position) {
    this.entitiesToLocalPositions.set(entity, position);
  }

  public getLocalPosition(entity: number): Position | undefined {
    return this.entitiesToLocalPositions.get(entity);
  }

  private updateAll() {
    this.entitiesToTargetPositions.forEach((targetPosition, entity) => {
      const localPos = this.entitiesToLocalPositions.get(entity);

      if (!localPos) {
        this.entitiesToLocalPositions.set(entity, targetPosition);
      } else {
        if (targetPosition) {
          if (Math.abs(targetPosition.x - localPos.x) < 0.1 && Math.abs(targetPosition.y - localPos.y) < 0.1) {
            return;
          }

          // Calculate new localPos values to get closer to targetPosition
          const newX = localPos.x + (targetPosition.x - localPos.x) / 6;
          const newY = localPos.y + (targetPosition.y - localPos.y) / 6;

          // Update localPos
          this.setLocalPosition(entity, { x: newX, y: newY });
        }
      }
    });
  }
}
