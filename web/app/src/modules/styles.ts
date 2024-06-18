export default {
    get custom() {
        return document.querySelector('#custom-styles')!.cloneNode();
    },
    *[Symbol.iterator]() {
        yield this.custom;
    },
};
