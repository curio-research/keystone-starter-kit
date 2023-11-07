import styled from '@emotion/styled';
import { Position } from 'core/schemas';
import { observer } from 'mobx-react';
import { useEffect, useState } from 'react';

export const tileSideWidth = 70;

interface PositionWrapperProps {
  position: Position;
}

export const PositionWrapper = styled.div<PositionWrapperProps>`
  position: absolute;
  left: ${(props) => props.position.x * tileSideWidth}px;
  bottom: ${(props) => props.position.y * tileSideWidth}px;
  width: ${tileSideWidth}px;
  height: ${tileSideWidth}px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
`;

interface ActivePositionWrapperProps {
  children: React.ReactNode;
  position: Position;
}

// Interpolates position values to make the movement smoother
export const ActivePositionWrapper = observer((props: ActivePositionWrapperProps) => {
  const { children, position } = props;
  const [localPos, setLocalPos] = useState<Position>(position);
  const [targetPosition, setTargetPosition] = useState<Position>(position);

  useEffect(() => {
    // console.log(toJS(position));
    setTargetPosition(position);
  }, [props]);

  useEffect(() => {
    const intervalId = setInterval(() => {
      // Check if localPos is close enough to targetPosition and stop the loop
      if (Math.abs(targetPosition.x - localPos.x) < 0.1 && Math.abs(targetPosition.y - localPos.y) < 0.1) {
        clearInterval(intervalId);
        return;
      }

      // Calculate new localPos values to get closer to targetPosition
      const newX = localPos.x + (targetPosition.x - localPos.x) / 6;
      const newY = localPos.y + (targetPosition.y - localPos.y) / 6;

      // Update localPos
      setLocalPos({ x: newX, y: newY });
    }, 10);

    // Clean up the interval when the component unmounts
    return () => clearInterval(intervalId);
  }, [targetPosition, localPos]);

  if (!localPos || !targetPosition) {
    return null;
  }

  return <PositionWrapper position={localPos}>{children}</PositionWrapper>;
});
