package helm

type HelmClient struct {
}

// func (h *HelmClient) ApplyManifest() error {
// 	opts := []client.PatchOption{client.ForceOwnership, client.FieldOwner(fieldOwnerOperator)}
// 	if err := h.client.Patch(context.TODO(), obj, client.Apply, opts...); err != nil {
// 		return fmt.Errorf("failed to update resource with server-side apply for obj %v: %v", objectStr, err)
// 	}
// 	return nil
// }

// func (h *HelmClient) RenderManifest() error {
// 	return nil
// }

// func (h *HelmClient) TransferManifestToObject() error {
// 	return nil
// }

// func (cl *Client) Patch(orig config.Config, patchFn config.PatchFunc) (string, error) {
// 	modified, patchType := patchFn(orig.DeepCopy())

// 	meta, err := patch(cl.client, orig, getObjectMetadata(orig), modified, getObjectMetadata(modified), patchType)
// 	if err != nil {
// 		return "", err
// 	}
// 	return meta.GetResourceVersion(), nil
// }
