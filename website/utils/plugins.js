const fs = require("fs");
const path = require("path");
const matter = require("gray-matter");
const title = require("title");
const pluginTitle = (plugin) => title(plugin.replace(/-/g, " "));


const tryReadPaths = (paths) => {
  for (const p of paths) {
    try {
      return fs.readFileSync(p);
    } catch (e) {
      if (p === paths[paths.length - 1]) {
        throw e;
      }
    }
  }
};

const getData = (plugin, type) => {
  try {
    const typeSingular = type === "sources" ? "source" : "destination";
    const oldsDocsPath = path.resolve(process.cwd(), `pages/docs/plugins/${type}/${plugin}/overview.md`);
    const newDocsPath = path.resolve(process.cwd(), `../plugins/${typeSingular}/${plugin}/docs/overview.md`);
    const overviewFile = tryReadPaths([oldsDocsPath, newDocsPath]);
    return {
      id: plugin,
      name: pluginTitle(plugin),
      stage: "Preview",
      ...matter(overviewFile).data,
    };
  } catch (e) {
    console.warn(`No overview file found for ${plugin}`);
    return { id: plugin, name: pluginTitle(plugin), stage: "Preview" };
  }
};

const getPluginsData = (type) => {
  const pluginsDir = path.resolve(process.cwd(), `pages/docs/plugins/${type}`);
  const files = fs.readdirSync(pluginsDir, { withFileTypes: true });
  const plugins = files
    .filter((file) => file.isDirectory())
    .map((file) => file.name);

  const withData = plugins.map((plugin) => [
    plugin,
    getData(plugin, type),
  ]);
  withData.sort((a, b) => a[1].name.localeCompare(b[1].name));
  return Object.fromEntries(withData);
};

module.exports = {
  getPluginsData,
};
