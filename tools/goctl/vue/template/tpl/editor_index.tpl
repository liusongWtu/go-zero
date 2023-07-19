<template>
    <{{.kebabObject}}-editor-form
        v-loading="loading"
        ref="refEditorForm"
        mode="page"
        :is-edit="true"
        :{{.lowerStartCamelObject}}="{{.lowerStartCamelObject}}Item"
        @submit="onSubmit"
    />
</template>

<script setup lang="ts">
    import { ref, computed, onMounted, onUnmounted } from 'vue';
    import { useRoute } from 'vue-router';
    import { useI18n } from 'vue-i18n';

    import { {{.upperStartCamelObject}}EditorForm, type {{.upperStartCamelObject}}EditorFormInstance } from './form';

    import { fetch{{.upperStartCamelObject}}Item, fetchUpdate{{.upperStartCamelObject}}, type {{.upperStartCamelObject}}Item, type ItemReq } from '@/api/{{.kebabObject}}';

    import { message } from '@/utils/message';

    defineOptions({ name: '{{.upperStartCamelObject}}EditorPage' });

    const refEditorForm = ref<{{.upperStartCamelObject}}EditorFormInstance>(null);

    const { t } = useI18n();
    const route = useRoute();

    const {{.lowerStartCamelObject}}Id = computed(() => Number(route.params.id));

    const loading = ref<boolean>(false);
    const {{.lowerStartCamelObject}}Item = ref<{{.upperStartCamelObject}}Item>({});

    const controller = new AbortController();

    const load{{.upperStartCamelObject}}ItemData = (id: number) => {
        if (loading.value) controller.abort();

        loading.value = true;
        const data: ItemReq = { id };
        fetch{{.upperStartCamelObject}}Item({ data, signal: controller.signal })
            .then(
                res => {
                    {{.lowerStartCamelObject}}Item.value = res.data;
                },
                err => {
                    if (!err?.isCancelRequest) {
                        message(t('common.dataLoadFailed', { error: String(err) }), { type: 'error' });
                    }
                }
            )
            .finally(() => {
                loading.value = false;
            });
    };

    onMounted(() => load{{.upperStartCamelObject}}ItemData({{.lowerStartCamelObject}}Id.value));

    onUnmounted(() => {
        if (loading.value) {
            controller.abort();
            loading.value = false;
        }
    });

    const onSubmit = () => {
        refEditorForm.value?.validate((valid, _fields) => {
            if (valid) {
                doUpdateConfig(refEditorForm.value.formData);
            }
        });
    };

    const doUpdate{{.upperStartCamelObject}} = (info: {{.upperStartCamelObject}}Item) => {
        const data = {
            id: {{.lowerStartCamelObject}}Id.value,
            name: info.name,
            data: info.data,
            code: info.code,
            type_id: info.type_id,
            type_name: info.type_name,
            type_code: info.type_code,
            schema: info.schema,
            remark: info.remark,
            status: info.status
        };
        fetchUpdate{{.upperStartCamelObject}}({ data }).then(
            () => {
                message(t('common.changeStatusMessage', { action: t('common.update'), content: info.name }), {
                    type: 'success'
                });
            },
            err => {
                message(
                    t('common.changeStatusFailed', {
                        action: t('common.update'),
                        content: info.name,
                        msg: err
                    }),
                    { type: 'error' }
                );
            }
        );
    };
</script>

<style lang="scss" scoped></style>
