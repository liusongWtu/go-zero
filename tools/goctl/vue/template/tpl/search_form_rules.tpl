import { useI18n } from 'vue-i18n';
import type { FormRules } from 'element-plus';
import { trimChar } from '@/utils/stringUtils';
import trim from 'ramda/es/trim';
import compose from 'ramda/es/compose';

const useRules = () => {
    const { t } = useI18n();

    const formRules: FormRules = {
        ids: [
            {
                trigger: 'blur',
                validator: (rule, value, callback) => {
                    if (!value) {
                        callback();
                    } else {
                        const v = compose(trimChar(','), trim)(value);
                        if (!v) {
                            callback();
                        } else {
                            const regex = new RegExp(/^\s*\d+\s*(,\s*\d+\s*)*\s*$/);
                            if (regex.test(v)) {
                                callback();
                            } else {
                                callback(new Error(t('{{.lowerStartCamelObject}}.search.idsRuleReg')));
                            }
                        }
                    }
                }
            }
        ]
    };
    return { formRules };
};

export { useRules };
