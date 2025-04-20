import { Component, ErrorInfo, ComponentType } from 'react';
import { SomethingWentWrong } from '@/pages/SomethingWentWrong';

interface State {
  hasError: boolean;
  error: Error | null;
}

export function ErrorWrapper<P>(
  WrappedComponent: ComponentType<P>,
  FallbackComponent: ComponentType = SomethingWentWrong,
): ComponentType<P> {
  return class ErrorBoundary extends Component<P, State> {
    public constructor(props: P) {
      super(props);
      this.state = { hasError: false, error: null };
    }

    public static getDerivedStateFromError(error: Error): State {
      return { hasError: true, error };
    }

    public override componentDidCatch(error: Error, errorInfo: ErrorInfo) {
      console.error('Error caught in ErrorBoundary:', error, errorInfo);
    }

    public override render() {
      if (this.state.hasError) {
        return <FallbackComponent />;
      }

      return <WrappedComponent {...this.props} />;
    }
  };
}
