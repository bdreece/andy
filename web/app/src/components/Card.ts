import styles from '../modules/styles';

const template = document.createElement('template');
template.innerHTML = `
<div class="card bg-base-200 shadow-xl">
    <div class="card-body">
        <h2 class="card-title">
            <slot name="title"></slot>
        </h2>
        
        <slot></slot>
        
        <div class="card-actions">
            <slot name="actions"></slot>
        </div>
    </div>
</div>
`;

export default class Card extends HTMLElement {
    constructor() {
        super();

        const shadow = this.attachShadow({ mode: 'open' });
        shadow.append(...styles, template.content.cloneNode(true));
    }
}

customElements.define('x-card', Card);

declare global {
    interface HTMLElementTagNameMap {
        'x-card': Card;
    }
}
