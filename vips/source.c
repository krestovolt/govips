#include "source.h"

VipsSourceCustom * create_go_custom_source()
{
	return vips_source_custom_new();
}

gulong connect_go_signal_read(VipsSourceCustom *source_custom, gpointer go_source)
{
	return g_signal_connect( source_custom, "read", G_CALLBACK(go_read), (gpointer) go_source );
}

gulong connect_go_signal_seek(VipsSourceCustom *source_custom, gpointer go_source)
{
	return g_signal_connect( source_custom, "seek", G_CALLBACK(go_seek), go_source );
}

void free_go_custom_source(VipsSourceCustom *source_custom, gulong rsig_handler_id, gulong ssig_handler_id)
{
	if (source_custom != NULL) {
		g_signal_handler_disconnect(source_custom, rsig_handler_id);
		g_signal_handler_disconnect(source_custom, ssig_handler_id);

		g_object_unref(source_custom);
	}
}

static gint64 go_read ( VipsSourceCustom *source_custom, void *buffer, gint64 length, gpointer go_source )
{
    return goSourceRead(go_source, buffer, length);
}

static gint64 go_seek ( VipsSourceCustom *source_custom, gint64 offset, int whence, gpointer go_source )
{
	return goSourceSeek(go_source, offset, whence);
}
