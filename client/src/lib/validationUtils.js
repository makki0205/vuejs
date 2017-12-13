import validator from 'validator';

const ValidationUtils = {

    isBlank(text) {
        return text === "" || text === null
    },

    isURL(url) {
        return validator.isURL(url)
    },

    pushMessage(message, vue) {
        vue.$toast.open({
            message,
            type: "is-danger"
        })
    }

}

export default ValidationUtils;