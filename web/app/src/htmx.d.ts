import type {
    HtmxConfigRequestDetail,
    HtmxValidateDetail,
    HtmxValidateFailedDetail,
    HtmxValidateHaltDetail,
} from './types/htmx';

declare global {
    interface HTMLElementEventMap {
        'htmx:configRequest': CustomEvent<HtmxConfigRequestDetail>;
        'htmx:validation:validate': CustomEvent<HtmxValidateDetail>;
        'htmx:validation:failed': CustomEvent<HtmxValidateFailedDetail>;
        'htmx:validation:halted': CustomEvent<HtmxValidateHaltDetail>;
    }
}

export {};
