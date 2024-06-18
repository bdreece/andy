export type HttpVerb =
    | 'get'
    | 'head'
    | 'post'
    | 'put'
    | 'delete'
    | 'connect'
    | 'options'
    | 'trace'
    | 'patch';

export type HtmxHeaderSpecification = Record<string, string>;

export type HtmxConfigRequestDetail = {
    boosted: boolean;
    useUrlParams: boolean;
    formData: FormData;
    parameters: Record<string, unknown>;
    unfilteredFormData: FormData;
    unfilteredParameters: Record<string, unknown>;
    headers: HtmxHeaderSpecification;
    target: Element;
    verb: HttpVerb;
    errors: HtmxValidateFailedDetail[];
    withCredentials: boolean;
    timeout: number;
    path: string;
    triggeringEvent: Event;
};

export type HtmxValidateDetail = {
    elt: Element;
};

export type HtmxValidateFailedDetail = {
    elt: Element;
    message: string;
    validity: ValidityState;
};

export type HtmxValidateHaltDetail = {
    elt: Element;
    errors: HtmxValidateFailedDetail[];
};
