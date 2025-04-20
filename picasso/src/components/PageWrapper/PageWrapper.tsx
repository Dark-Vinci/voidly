import { ReactNode } from 'react';
import { SafeAreaView, StyleSheet, ScrollView } from 'react-native';

import { fi } from '@utils';

interface props {
  children: ReactNode;
  backgroundColor?: string;
}

export function PageWrapper({ children, backgroundColor }: props): JSX.Element {
  return (
    <SafeAreaView
      style={[style.container, fi(!!backgroundColor, { backgroundColor })]}
    >
      <ScrollView
        style={{
          flex: 1,
          width: '100%',
          height: '100%',
          backgroundColor: 'white',
          flexGrow: 1,
        }}
        contentContainerStyle={{
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        {children}
      </ScrollView>
    </SafeAreaView>
  );
}

const style = StyleSheet.create({
  container: {
    backgroundColor: 'blue',
  },
});
