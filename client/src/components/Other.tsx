import styled from '@emotion/styled';
import { Position } from 'core/schemas';
import { positionWrapperState } from 'index';
import { observer } from 'mobx-react';
import { useEffect } from 'react';

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
  entity: number;
}

// Interpolates position values to make the movement smoother
export const ActivePositionWrapper = observer((props: ActivePositionWrapperProps) => {
  const { children, position, entity } = props;

  useEffect(() => {
    positionWrapperState.setTargetPosition(entity, position);
  }, [position]);

  const localPosition = positionWrapperState.getLocalPosition(entity);

  if (!localPosition) {
    return null;
  }

  return <PositionWrapper position={localPosition}>{children}</PositionWrapper>;
});
