import { SafeAreaView, Text, View } from 'react-native';
import { JSX, useState } from 'react';
import * as Font from 'expo-font';
import AppLoading from 'expo-app-loading';
import { ErrorWrapper, If, IfElse } from '@components';

interface props {
  name: string;
}

function Wrapped({ name }: props): JSX.Element {
  return (
    <View>
      <If condition={true} element={<Text>PANICKER {name}</Text>} />

      <IfElse
        condition={true}
        elseElement={<Text>painstaking elemination</Text>}
        ifElement={<Text>PANICKER</Text>}
      />
    </View>
  );
}

const MeApp = ErrorWrapper(Wrapped);

export default function App(): JSX.Element {
  const [fontsLoaded, setFontsLoaded] = useState(false);

  const loadFonts = async () => {
    await Font.loadAsync({
      MyFontBold: require('./assets/fonts/Junction-bold.otf'),
      MyFontLight: require('./assets/fonts/Junction-light.otf'),
      MyFontRegular: require('./assets/fonts/Junction-regular.otf'),
    });
  };

  if (!fontsLoaded) {
    return (
      <AppLoading
        startAsync={loadFonts}
        onFinish={() => setFontsLoaded(true)}
        onError={console.warn}
      />
    );
  }

  return (
    <SafeAreaView>
      <View>
        <If
          condition={true}
          element={
            <Text style={{ fontFamily: 'MyFontBold' }}>
              What is what is this going to do your name
            </Text>
          }
        />

        <Text>What is what is this going to do your name</Text>
        <MeApp name={'melon'} />

        <IfElse
          condition={true}
          elseElement={<Text>painstaking elemination</Text>}
          ifElement={<Text>Hello world</Text>}
        />
      </View>
    </SafeAreaView>
  );
}
