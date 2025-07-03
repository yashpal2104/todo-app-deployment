import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    // Consult https://kit.svelte.dev/docs/integrations#preprocessors
    // for more information about preprocessors
    preprocess: vitePreprocess(),

    kit: {
        // This is the crucial part.
        // We are replacing the default 'adapter-auto' with 'adapter-static'.
        adapter: adapter({
            // default options are fine, but we can be explicit
            pages: 'build',      // The output directory for pages
            assets: 'build',     // The output directory for assets
            fallback: 'index.html', // Create a fallback for single-page app behavior
            precompress: false
        })
    }
};

export default config;
