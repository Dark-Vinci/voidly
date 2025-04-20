import { ReactNode } from 'react';

interface ifElseProps {
  condition: boolean;
  ifElement: ReactNode;
  elseElement: ReactNode;
}

export function IfElse({
  condition,
  ifElement,
  elseElement,
}: ifElseProps): JSX.Element {
  if (condition) {
    return ifElement as JSX.Element;
  }

  return elseElement as JSX.Element;
}
