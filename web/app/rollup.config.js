import nodeResolve from '@rollup/plugin-node-resolve'
import typescript from '@rollup/plugin-typescript'
import json from '@rollup/plugin-json';
import postcss from 'rollup-plugin-postcss';
import autoprefixer from 'autoprefixer';
import cssnano from 'cssnano';
import tailwindcss from 'tailwindcss';

/** @type {import('rollup').RollupOptions} */
export default {
    input: 'src/index.ts',
    output: {
        dir: 'dist',
        format: 'esm',
    },
    plugins: [
        nodeResolve(),
        typescript(),
        json(),
        postcss({
            extract: true,
            plugins: [
                autoprefixer(),
                cssnano(),
                tailwindcss(),
            ],
        }),
    ]
}
