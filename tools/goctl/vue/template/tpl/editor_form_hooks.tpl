import { ref, watch, createVNode, type Ref } from 'vue';
import { useI18n } from 'vue-i18n';

import { ElScrollbar, type FormInstance, type FormItemProp, type FormValidateCallback } from 'element-plus';

import { JsonCodeMirror } from '@/components/JsonCodeMirror';
import { addDialog } from '@/components/ReDialog';
import { CustomDialogHF } from '@/common/CustomDialogHF';

import equals from 'ramda/es/equals';

import { type CommonStatus } from '@/defines';
import { type {{.upperStartCamelObject}}Item } from '@/api/{{.kebabObject}}';

export interface {{.upperStartCamelObject}}EditorFormProps {
    mode: 'page' | 'dialog';
    hasSubmitButton?: boolean;
    isEdit: boolean;
    {{.lowerStartCamelObject}}: {{.upperStartCamelObject}}Item;
}

export function useForm(refForm: Ref<FormInstance>, props: {{.upperStartCamelObject}}EditorFormProps, emit: (evt: 'submit') => void) {
    const { t } = useI18n();

    const formData = ref<{{.upperStartCamelObject}}Item>(props.{{.lowerStartCamelObject}});


    const onChangeStatus = (_val: CommonStatus): void => {
        //TODO: nothing
    };

    const onClickedSubmit = () => emit('submit');

    const validate = (callback?: FormValidateCallback) => {
        return refForm.value?.validate(callback);
    };

    const validateField = (props: FormItemProp, callback?: FormValidateCallback) => {
        return refForm.value?.validateField(props, callback);
    };

    watch(
        () => props.{{.lowerStartCamelObject}},
        value => {
            if (!equals(value, formData.value)) {
                formData.value = value;
            }
        }
    );

    return { formData, onChangeStatus, onClickedSubmit, validate, validateField };
}
