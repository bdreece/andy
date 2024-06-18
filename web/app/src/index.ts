import * as iconify from 'iconify-icon';
import { icons as tabler } from '@iconify-json/tabler';
import 'htmx.org';

import * as csrf from './modules/csrf';
import './index.css';

iconify.addCollection(tabler);

document.body.addEventListener('htmx:configRequest', csrf.onConfigRequest);

export * as components from './components';
