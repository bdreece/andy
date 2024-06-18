import styles from '../modules/styles';

const template = document.createElement('template');
template.innerHTML = `
<div role="alert" class="alert">
    <slot></slot>
</div>
`;

export default class Alert extends HTMLElement {
    #alert;

    constructor() {
        super();

        const shadow = this.attachShadow({ mode: 'closed' });
        shadow.append(...styles, template.content.cloneNode(true));
        this.#alert = shadow.querySelector('div')!;
    }

    connectedCallback() {
        const icon = this.getAttribute('icon');
        if (icon) {
            const iconifyIcon = document.createElement('iconify-icon');
            iconifyIcon.icon = icon;

            this.#alert.insertAdjacentElement('afterbegin', iconifyIcon);
        }

        const variant = this.getAttribute('variant');
        switch (variant) {
            case 'info':
                this.#alert.classList.add('alert-info');
                break;
            case 'success':
                this.#alert.classList.add('alert-success');
                break;
            case 'warning':
                this.#alert.classList.add('alert-warning');
                break;
            case 'error':
                this.#alert.classList.add('alert-error');
                break;
            default:
                break;
        }
    }
}

customElements.define('x-alert', Alert);

declare global {
    interface HTMLElementTagNameMap {
        'x-alert': Alert;
    }
}
