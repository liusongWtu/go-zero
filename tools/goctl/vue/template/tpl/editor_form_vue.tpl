<template>
    <el-scrollbar class="!px-[20px]" :max-height="mode == 'dialog' ? '300px' : undefined">
        <el-form
            ref="refForm"
            class="!my-[30px]"
            :model="formData"
            :rules="formRules"
            label-width="100px"
            size="small"
            label-position="left"
        >
{{.editorFormVueFields}}            
        </el-form>
        <el-button v-if="mode == 'page'" class="w-full" type="primary" @click="onClickedSubmit">{{
            printf "$t('common.submit')"
        }}</el-button>
    </el-scrollbar>
</template>

<script setup lang="ts">
    import { ref } from 'vue';
    import { type FormInstance } from 'element-plus';
    import { CommonStatusSelector } from '@/common/CommonStatusSelector';

    import { useForm, type {{.upperStartCamelObject}}EditorFormProps } from './formHooks';
    import { useRules } from './formRules';

    import { CommonStatusActivate } from '@/defines';
    import { type {{.upperStartCamelObject}}Item } from '@/api/{{.kebabObject}}';

    defineOptions({ name: '{{.upperStartCamelObject}}EditorForm' });

    const props = withDefaults(defineProps<{{.upperStartCamelObject}}EditorFormProps>(), {
        mode: 'dialog',
        isEdit: false,
        config: (): {{.upperStartCamelObject}}Item => ({ status: CommonStatusActivate })
    });

    const emit = defineEmits<{
        (evt: 'submit'): void;
    }>();

    const refForm = ref<FormInstance>();

    const {
        formData,
        dataErrors,
        onChangeSchemaId,
        onClickedSchemaInput,
        onClickedDataInput,
        onChangeStatus,
        onClickedSubmit,
        validate,
        validateField
    } = useForm(refForm, props, emit);

    const formRules = useRules(dataErrors).formRules;

    defineExpose({ formData, validate, validateField });
</script>

<style lang="scss"></style>