import { reactive } from 'vue';
import { useI18n } from 'vue-i18n';
import type { FormRules } from 'element-plus';


const useRules = function () {
    const { t } = useI18n();

    const formRules = reactive<FormRules>({
        status: [{ required: true, message: t('{{.lowerStartCamelObject}}.{{.lowerStartCamelObject}}Editor.statusReg'), trigger: 'allow' }]
    });

    return { formRules };
};

export { useRules };
