import styles from '../modules/styles';

const template = document.createElement('template');
template.innerHTML = `
    <iconify-icon icon="tabler:key"></iconify-icon>
    
    <input
        class="grow"
        type="password"
        placeholder="Password"
    >
`;

export default class PasswordInput extends HTMLElement {
    static formAssociated = true;

    #internals;
    #inner;

    constructor() {
        super();

        this.#internals = this.attachInternals();
        const shadow = this.attachShadow({
            mode: 'closed',
            delegatesFocus: true,
        });

        shadow.append(...styles, template.content.cloneNode(true));
        this.#inner = shadow.querySelector('input')!;
        this.value = '';
    }

    get value() {
        return this.#inner.value;
    }
    set value(value) {
        this.#inner.value = value;
        this.#internals.setFormValue(value);
    }

    get form() {
        return this.#internals.form;
    }
    get name() {
        return this.getAttribute('name');
    }
    get type() {
        return 'password';
    }
    get validity() {
        return this.#internals.validity;
    }
    get validationMessage() {
        return this.#internals.validationMessage;
    }
    get willValidate() {
        return this.#internals.willValidate;
    }

    checkValidity() {
        return this.#internals.checkValidity();
    }
    reportValidity() {
        return this.#internals.reportValidity();
    }

    connectedCallback() {}
}

customElements.define('x-password-input', PasswordInput);

declare global {
    interface HTMLElementTagNameMap {
        'x-password-input': PasswordInput;
    }
}
