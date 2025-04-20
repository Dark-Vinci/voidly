module.exports = function (api) {
  api.cache(true);
  
  return {
    presets: ['babel-preset-expo'],
    plugins: [
      [
        'module-resolver',
        {
          alias: {
            '@components': './src/components',
            '@containers': './src/containers',
            "@pages": "./src/pages",
            "@state": "./src/state",
            "@types": "./src/types",
            "@": "./src",
            "@navigations": "./src/navigations",
            "@utils": "./src/utils"
          },
        },
      ],
    ],
  };
};