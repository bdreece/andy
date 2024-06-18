import type { HtmxConfigRequest } from '../types/htmx';

import cookies from './cookies';

export function onConfigRequest(e: CustomEvent<HtmxConfigRequest>) {
    e.detail.headers['X-CSRF-Token'] = cookies.get('_andy-csrf') ?? '';
}
